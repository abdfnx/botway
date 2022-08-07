project = "botway-cdn"

app "botway-cdn" {
    labels = {
        "service" = "botway-cdn",
        "env" = "dev"
    }

    build {
        use "docker" {}
    }

    deploy {
        use "docker" {}
    }
}
