require 'osmosis/instance'
module Osmosis
  class Manager
    attr_reader :socket, :error

    def initialize(app)
      @app = app
      connect!
    end

    def call(env)
      env['osmosis'] = Osmosis::Instance.new(socket)
      @app.call(env)
    end

    private

    def connect!
      begin
        @socket ||= UNIXSocket.new(path)
      rescue Exception => e
        @error = e
      end
    end

    def connected?
      error.nil?
    end

    def path
      tmp = Rails.application.paths['tmp'].existent.first
      tmp + '/sockets/osmosis.sock'
    end
  end
end
