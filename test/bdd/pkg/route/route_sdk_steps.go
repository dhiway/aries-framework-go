/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package route

import (
	"fmt"

	"github.com/DATA-DOG/godog"

	"github.com/hyperledger/aries-framework-go/pkg/client/route"
	"github.com/hyperledger/aries-framework-go/test/bdd/pkg/context"
)

// SDKSteps is steps for route using client SDK
type SDKSteps struct {
	bddContext *context.BDDContext
}

// NewRouteSDKSteps return steps for router using client SDK
func NewRouteSDKSteps(ctx *context.BDDContext) *SDKSteps {
	return &SDKSteps{
		bddContext: ctx,
	}
}

// CreateRouteClient creates route client
func (d *SDKSteps) CreateRouteClient(agentID string) error {
	// create new route client
	routeClient, err := route.New(d.bddContext.AgentCtx[agentID])
	if err != nil {
		return fmt.Errorf("failed to create new route client: %w", err)
	}

	d.bddContext.RouteClients[agentID] = routeClient

	return nil
}

// RegisterRoute registers the router for the agent
func (d *SDKSteps) RegisterRoute(agentID, varName string) error {
	err := d.bddContext.RouteClients[agentID].Register(d.bddContext.Args[varName])
	if err != nil {
		return fmt.Errorf("failed to handle invitation: %w", err)
	}

	return nil
}

// RegisterSteps registers router steps
func (d *SDKSteps) RegisterSteps(s *godog.Suite) {
	s.Step(`^"([^"]*)" creates a route exchange client$`, d.CreateRouteClient)
	s.Step(`^"([^"]*)" sets "([^"]*)" as the router$`, d.RegisterRoute)
}
