module "discovery" {
  source      = "../modules/discovery"
  environment = "production"
}

module "darkstat" {
  source         = "../modules/darkstat"
  environment    = "production"
  tag            = "production-arm64"
  discovery_path = module.discovery.discovery_path
  ipv4_address   = module.data_cluster.node_darklab.ipv4_address
  enable_restarts = true

  RELAY_HOST         = "https://darkrelay.dd84ai.com"
  SITE_ROOT          = "/fl-data-discovery/"
  FLDARKSTAT_HEADING = <<-EOT
  <a href="https://github.com/darklab8/fl-darkstat">Darkstat</a> from <a href="https://darklab8.github.io/blog/pet_projects.html#Freelancercommunity">DarkTools</a> for <a href="https://github.com/darklab8/fl-data-discovery">Freelancer Discovery</a>
  EOT

  stat_prefix  = "darkstat"
  relay_prefix = "darkrelay"
  rpc_prefix   = "darkgrpc"
  zone         = "dd84ai.com"
}

resource "random_string" "random_password" {
  length  = 32
  special = false
}
resource "random_string" "random_secret" {
  length  = 32
  special = false
}

module "discovery_dev" {
  source      = "../modules/discovery_dev"
  environment = "dev"
}

module "darkstat_dev" {
  source         = "../modules/darkstat"
  environment    = "dev"
  tag            = "production-arm64"
  discovery_path = module.discovery_dev.discovery_path
  ipv4_address   = module.data_cluster.node_darklab.ipv4_address

  SITE_ROOT          = "/"
  FLDARKSTAT_HEADING = <<-EOT
  <a href="https://github.com/darklab8/fl-darkstat">Darkstat</a> from <a href="https://darklab8.github.io/blog/pet_projects.html#Freelancercommunity">DarkTools</a> for <a href="https://github.com/darklab8/fl-data-discovery">Freelancer Discovery</a>
  EOT

  stat_prefix  = "darkstat-dev"
  relay_prefix = "darkrelay-dev"
  zone         = "dd84ai.com"
  enable_restarts = true

  password = random_string.random_password.result
  secret   = random_string.random_secret.result
}

module "vanilla" {
  source      = "../modules/vanilla"
  environment = "production"
}

module "darkstat_vanilla" {
  source         = "../modules/darkstat"
  environment    = "vanilla"
  tag            = "production-arm64"
  discovery_path = module.vanilla.vanilla_path
  ipv4_address   = module.data_cluster.node_darklab.ipv4_address

  SITE_ROOT          = "/fl-data-discovery/"
  FLDARKSTAT_HEADING = <<-EOT
  <a href="https://github.com/darklab8/fl-darkstat">Darkstat</a> from <a href="https://darklab8.github.io/blog/pet_projects.html#Freelancercommunity">DarkTools</a> for Freelancer Vanilla
  EOT

  stat_prefix  = "darkstat-vanilla"
  rpc_prefix   = "darkgrpc-vanilla"
  zone         = "dd84ai.com"
  enable_restarts = false
}
