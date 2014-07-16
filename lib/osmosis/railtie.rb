require 'osmosis/manager'

module Osmosis
  class Railtie < ::Rails::Railtie
    initializer 'osmosis.connect' do |app|
      app.middleware.use Osmosis::Manager
    end
  end
end
