require 'osmosis/instance'

module Osmosis
  class Middleware

    def initialize(app)
      @app = app
    end

    def call(env)
      env['osmosis'] = Osmosis::Instance.new(Rails.configuration.osmosis.backend.channel)
      @app.call(env)
    end
  end
end
