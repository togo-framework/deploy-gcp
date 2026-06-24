// Package gcp deploys a togo app to Google Cloud (Cloud Run) by driving the `gcloud` CLI.
// Select with deploy.provider=gcp; requires the gcloud CLI authenticated.
package gcp

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/togo-framework/deploy"
	"github.com/togo-framework/togo"
)

func init() { deploy.RegisterDriver("gcp", New) }

// New checks the gcloud CLI is present.
func New(_ *togo.Kernel) (deploy.Deployer, error) {
	if _, err := exec.LookPath("gcloud"); err != nil {
		return nil, fmt.Errorf("deploy-gcp: the %q CLI is required (install + authenticate it)", "gcloud")
	}
	return &driver{}, nil
}

type driver struct{}

func run(ctx context.Context, name string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	return string(out), err
}

func region(s deploy.Spec, def string) string {
	if s.Region != "" {
		return s.Region
	}
	return def
}

func (d *driver) Provision(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	return d.Deploy(ctx, spec)
}

func (d *driver) Deploy(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	out, err := run(ctx, "gcloud", "run", "deploy", spec.App, "--image", spec.Image, "--region", region(spec, "us-central1"), "--platform", "managed", "--allow-unauthenticated", "--quiet", "--format", "value(status.url)")
	if err != nil {
		return nil, fmt.Errorf("gcloud run deploy: %v: %s", err, out)
	}
	return &deploy.Result{URL: out, Message: "deployed to Cloud Run", Raw: map[string]any{"out": out}}, nil
}

func (d *driver) Destroy(ctx context.Context, spec deploy.Spec) error {
	_, err := run(ctx, "gcloud", "run", "services", "delete", spec.App, "--region", region(spec, "us-central1"), "--quiet")
	return err
}

func (d *driver) Status(ctx context.Context, spec deploy.Spec) (*deploy.Status, error) {
	out, err := run(ctx, "gcloud", "run", "services", "describe", spec.App, "--region", region(spec, "us-central1"), "--format", "value(status.conditions[0].status)")
	return &deploy.Status{Healthy: err == nil, Detail: out}, err
}
