/*
Copyright 2021 Adobe. All rights reserved.
This file is licensed to you under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License. You may obtain a copy
of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
OF ANY KIND, either express or implied. See the License for the specific language
governing permissions and limitations under the License.
*/

package v2

import (
	"github.com/adobe/cluster-registry/pkg/apiserver/models"
	"net/http"
	"strconv"

	registryv1 "github.com/adobe/cluster-registry/pkg/api/registry/v1"
	"github.com/adobe/cluster-registry/pkg/apiserver/errors"
	"github.com/adobe/cluster-registry/pkg/apiserver/web"
	"github.com/adobe/cluster-registry/pkg/auth"
	"github.com/adobe/cluster-registry/pkg/config"
	"github.com/adobe/cluster-registry/pkg/database"
	monitoring "github.com/adobe/cluster-registry/pkg/monitoring/apiserver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Handler interface
type Handler interface {
	GetCluster(echo.Context) error
	ListClusters(echo.Context) error
	Register(*echo.Group)
}

// handler struct
type handler struct {
	db        database.Db
	appConfig *config.AppConfig
	metrics   monitoring.MetricsI
}

// NewHandler func
func NewHandler(appConfig *config.AppConfig, d database.Db, m monitoring.MetricsI) Handler {
	h := &handler{
		db:        d,
		metrics:   m,
		appConfig: appConfig,
	}
	return h
}

func (h *handler) Register(v2 *echo.Group) {
	a, err := auth.NewAuthenticator(h.appConfig, h.metrics)
	if err != nil {
		log.Fatalf("Failed to initialize authenticator: %v", err)
	}
	clusters := v2.Group("/clusters", a.VerifyToken(), web.RateLimiter(h.appConfig))
	clusters.GET("/:name", h.GetCluster)
	clusters.GET("", h.ListClusters)
}

// GetCluster godoc
// @Summary Get an cluster
// @Description Get an cluster. Auth is required
// @ID v2-get-cluster
// @Tags cluster
// @Accept  json
// @Produce  json
// @Param name path string true "Name of the cluster to get"
// @Success 200 {object} registryv1.ClusterSpec
// @Failure 400 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Security bearerAuth
// @Router /v2/clusters/{name} [get]
func (h *handler) GetCluster(ctx echo.Context) error {

	name := ctx.Param("name")
	c, err := getCluster(h.db, name)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.NewError(err))
	}

	if c == nil {
		return ctx.JSON(http.StatusNotFound, errors.NotFound())
	}

	return ctx.JSON(http.StatusOK, newClusterResponse(ctx, c))
}

// ListClusters
// @Summary List clusters
// @Description List all or a subset of clusters. Use conditions to filter clusters based on their fields.
// @ID v2-get-clusters
// @Tags cluster
// @Accept  json
// @Produce  json
// @Param conditions query []string false "Filter conditions" collectionFormat(multi)
// @Param offset query integer false "Offset to start pagination search results (default is 0)"
// @Param limit query integer false "The number of results per page (default is 200)"
// @Success 200 {object} clusterList
// @Failure 500 {object} errors.Error
// @Security bearerAuth
// @Router /v2/clusters [get]
func (h *handler) ListClusters(ctx echo.Context) error {
	var clusters []registryv1.Cluster
	var count int

	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		limit = 200
	}

	filter := database.NewDynamoDBFilter()
	queryConditions := getQueryConditions(ctx)

	if len(queryConditions) == 0 {
		clusters, count, more, _ := h.db.ListClusters(offset, limit, "", "", "", "")
		return ctx.JSON(http.StatusOK, newClusterListResponse(clusters, count, offset, limit, more))
	}

	for _, qc := range queryConditions {
		condition, err := models.NewFilterConditionFromQuery(qc)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, errors.NewError(err))
		}
		filter.AddCondition(condition)
	}

	clusters, count, more, _ := h.db.ListClustersWithFilter(offset, limit, filter)
	return ctx.JSON(http.StatusOK, newClusterListResponse(clusters, count, offset, limit, more))
}

// getCluster by standard name or short name
func getCluster(db database.Db, name string) (*registryv1.Cluster, error) {

	var c *registryv1.Cluster
	var err error

	c, err = db.GetCluster(name)
	if err != nil {
		return nil, err
	}

	if c == nil {
		dashName, err := web.GetClusterDashName(name)
		if err != nil {
			log.Infof("Cluster %s is not a short name. Error: %v", name, err.Error())
		} else {
			c, err = db.GetCluster(dashName)
			if err != nil {
				return nil, err
			}
		}
	}
	return c, nil
}

func getQueryConditions(ctx echo.Context) []string {
	for k, v := range ctx.QueryParams() {
		if k == "conditions" {
			return v
		}
	}
	return []string{}
}
