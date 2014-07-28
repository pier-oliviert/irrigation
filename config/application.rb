require File.expand_path('../boot', __FILE__)

require 'rails/all'

Bundler.require(*Rails.groups)

module Irrigation
  class Application < Rails::Application
    paths.add "config/locales",      glob: "**/*.{rb,yml}"
    paths.add "sockets/osmosis",     with: "tmp/sockets/osmosis.sock"

    paths.add 'lib/osmosis.rb', eager_load: true
    # One day, paths won't be totally retarded
    require paths['lib/osmosis.rb'].existent.first

    config.time_zone = 'Eastern Time (US & Canada)'
    config.i18n.default_locale = :'fr-CA'
  end
end
