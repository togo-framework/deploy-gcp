# deploy-gcp — docs

**Google Cloud deploy.** Deploy to Cloud Run via the authenticated `gcloud` CLI.

## Install

```bash
togo install togo-framework/deploy-gcp
```

Registers on the [`deploy`](https://github.com/togo-framework/deploy) base; select it with **deploy.provider in togo.yaml (or DEPLOY_PROVIDER)**, then use **`togo deploy`**.

## Interface

`Deployer` — `Provision`/`Deploy`/`Destroy`/`Status` over a `Spec{App,Dir,BuildCmd,Host,User,Image,Region,Domain}` built from your `togo.yaml`.

## Usage & notes

Requires `gcloud` installed + authenticated (project set). Deploys `spec.Image` to Cloud Run.

## Example

```bash
togo deploy --provider gcp --dry-run   # preview the plan
togo deploy --provider gcp
```

## Links

- [Cloud Run](https://cloud.google.com/run/docs)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/deploy-gcp)
