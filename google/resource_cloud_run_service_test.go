package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestCloudrunAnnotationDiffSuppress(t *testing.T) {
	cases := map[string]struct {
		K, Old, New        string
		ExpectDiffSuppress bool
	}{
		"missing run.googleapis.com/operation-id": {
			K:                  "metadata.0.annotations.run.googleapis.com/operation-id",
			Old:                "12345abc",
			New:                "",
			ExpectDiffSuppress: true,
		},
		"missing run.googleapis.com/ingress": {
			K:                  "metadata.0.annotations.run.googleapis.com/ingress",
			Old:                "all",
			New:                "",
			ExpectDiffSuppress: true,
		},
		"explicit run.googleapis.com/ingress": {
			K:                  "metadata.0.annotations.run.googleapis.com/ingress",
			Old:                "all",
			New:                "internal",
			ExpectDiffSuppress: false,
		},
	}
	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {
			if got := cloudrunAnnotationDiffSuppress(tc.K, tc.Old, tc.New, nil); got != tc.ExpectDiffSuppress {
				t.Errorf("got %t; want %t", got, tc.ExpectDiffSuppress)
			}
		})
	}
}

func TestAccCloudRunService_cloudRunServiceUpdate(t *testing.T) {
	t.Parallel()

	project := acctest.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + RandString(t, 6)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "50", "300"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

// this test checks that Terraform does not fail with a 409 recreating the same service
func TestAccCloudRunService_foregroundDeletion(t *testing.T) {
	t.Parallel()

	project := acctest.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + RandString(t, 6)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: " ", // very explicitly add a space, as the test runner fails if this is just ""
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdate(name, project, "10", "600"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceUpdate(name, project, concurrency, timeoutSeconds string) string {
	return fmt.Sprintf(`
resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  metadata {
  namespace = "%s"
  annotations = {
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        ports {
          container_port = 8080
        }
      }
      container_concurrency = %s
      timeout_seconds = %s
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
    tag             = "magic-module"
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }
}
`, name, project, concurrency, timeoutSeconds)
}

func TestAccCloudRunService_secretVolume(t *testing.T) {
	t.Parallel()

	project := acctest.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + RandString(t, 6)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretVolume(name, project, "secret-"+RandString(t, 5), "secret-"+RandString(t, 6), "google_secret_manager_secret.secret1.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretVolume(name, project, "secret-"+RandString(t, 10), "secret-"+RandString(t, 11), "google_secret_manager_secret.secret2.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceUpdateWithSecretVolume(name, project, secretName1, secretName2, secretRef string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret1" {
  secret_id = "%s"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret" "secret2" {
  secret_id = "%s"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret1-version-data" {
  secret = google_secret_manager_secret.secret1.name
  secret_data = "secret-data1"
}

resource "google_secret_manager_secret_version" "secret2-version-data" {
  secret = google_secret_manager_secret.secret2.name
  secret_data = "secret-data2"
}

resource "google_secret_manager_secret_iam_member" "secret1-access" {
  secret_id = google_secret_manager_secret.secret1.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret1]
}

resource "google_secret_manager_secret_iam_member" "secret2-access" {
  secret_id = google_secret_manager_secret.secret2.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret2]
}

resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        volume_mounts {
          name = "a-volume"
          mount_path = "/secrets"
        }
      }
      volumes {
        name = "a-volume"
        secret {
          secret_name = %s
          items {
            key = "1"
            path = "my-secret"
          }
        }
      }
    }
  }

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret1-version-data, google_secret_manager_secret_version.secret2-version-data]
}
`, secretName1, secretName2, name, secretRef, project)
}

func TestAccCloudRunService_secretEnvironmentVariable(t *testing.T) {
	t.Parallel()

	project := acctest.GetTestProjectFromEnv()
	name := "tftest-cloudrun-" + RandString(t, 6)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretEnvVar(name, project, "secret-"+RandString(t, 5), "secret-"+RandString(t, 6), "google_secret_manager_secret.secret1.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
			{
				Config: testAccCloudRunService_cloudRunServiceUpdateWithSecretEnvVar(name, project, "secret-"+RandString(t, 10), "secret-"+RandString(t, 11), "google_secret_manager_secret.secret2.secret_id"),
			},
			{
				ResourceName:            "google_cloud_run_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata.0.resource_version", "status.0.conditions"},
			},
		},
	})
}

func testAccCloudRunService_cloudRunServiceUpdateWithSecretEnvVar(name, project, secretName1, secretName2, secretRef string) string {
	return fmt.Sprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret1" {
  secret_id = "%s"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret" "secret2" {
  secret_id = "%s"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret1-version-data" {
  secret = google_secret_manager_secret.secret1.name
  secret_data = "secret-data1"
}

resource "google_secret_manager_secret_version" "secret2-version-data" {
  secret = google_secret_manager_secret.secret2.name
  secret_data = "secret-data2"
}

resource "google_secret_manager_secret_iam_member" "secret1-access" {
  secret_id = google_secret_manager_secret.secret1.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret1]
}

resource "google_secret_manager_secret_iam_member" "secret2-access" {
  secret_id = google_secret_manager_secret.secret2.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret2]
}

resource "google_cloud_run_service" "default" {
  name     = "%s"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/cloudrun/hello"
        env {
          name = "SECRET_ENV_VAR"
          value_from {
            secret_key_ref {
              name = %s
              key = "1"
            }
          }
        }
      }
    }
  }

  metadata {
    namespace = "%s"
    annotations = {
      generated-by = "magic-modules"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret1-version-data, google_secret_manager_secret_version.secret2-version-data]
}
`, secretName1, secretName2, name, secretRef, project)
}
