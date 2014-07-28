require 'osmosis/manager'
require 'osmosis/middleware'

module Osmosis
  class Railtie < ::Rails::Railtie

    initializer 'osmosis.configure' do |app|
      app.config.osmosis = Osmosis::Manager.new
    end

    initializer 'osmosis.start' do |app|
      app.config.osmosis.start!
    end

    initializer 'osmosis.manager' do |app|
      app.middleware.use Osmosis::Middleware
    end
  end
end
