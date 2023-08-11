import * as pulumi from "@pulumi/pulumi";
import * as gcp from "@pulumi/gcp";

const serviceName =
  new pulumi.Config("service").get("name") || pulumi.getProject();
const location = gcp.config.region || "asia-east1";

const github_connection = gcp.cloudbuildv2.Connection.get(
  "github-connection",
  pulumi.interpolate`projects/${gcp.config.project}/locations/${location}/connections/GitHub`
);

const repository = new gcp.cloudbuildv2.Repository("repository", {
  location,
  name: serviceName,
  parentConnection: github_connection.name,
  remoteUri: pulumi.interpolate`https://github.com/ride-app/${serviceName}.git`,
});

new gcp.cloudbuild.Trigger("build-trigger", {
  name: serviceName,
  location,
  repositoryEventConfig: {
    repository: repository.id,
    push: {
      branch: "^main$",
    },
  },
  substitutions: {
    _LOG_DEBUG: new pulumi.Config().get("logDebug") ?? "false",
  },
  filename: "cloudbuild.yaml",
  includeBuildLogs: "INCLUDE_BUILD_LOGS_WITH_STATUS",
});

// const imageconfig = new pulumi.Config("image");

// // Cloud Run
// const service = new gcp.cloudrun.Service("service", {
//   name: serviceName,
//   location,
//   template: {
//     metadata: {
//       annotations: {
//         "autoscaling.knative.dev/maxScale": "10",
//       },
//     },
//     spec: {
//       containers: [
//         {
//           image: `asia-south2-docker.pkg.dev/${
//             gcp.config.project
//           }/docker-registry/${serviceName}:${
//             imageconfig.get("tag") ?? "latest"
//           }`,
//           ports: [{ containerPort: 50051, name: "h2c" }],
//           envs: [
//             {
//               name: "FIREBASE_PROJECT_ID",
//               value: gcp.config.project,
//             },
//             {
//               name: "DEBUG",
//               value: new pulumi.Config().get("debug") ?? "false",
//             },
//           ],
//         },
//       ],
//     },
//   },
// });

// const policyData = gcp.organizations.getIAMPolicy({
//   bindings: [
//     {
//       role: "roles/run.invoker",
//       members: ["allUsers"],
//     },
//   ],
// });

// const noauthIamPolicy = new gcp.cloudrun.IamPolicy("no-auth-iam-policy", {
//   location: service.location,
//   project: service.project,
//   service: service.name,
//   policyData: policyData.then((noauthIAMPolicy) => noauthIAMPolicy.policyData),
// });

// export const name = service.name;
