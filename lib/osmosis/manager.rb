require 'osmosis/backend'
require 'osmosis/clients'

module Osmosis
  class Manager
    attr_reader :backend, :clients

    def initialize
      @backend = Osmosis::Backend.new
      @clients = Osmosis::Clients.new
    end

    def start!
      @clients.start!

      @backend.start! do |msg|
        @clients.channel.push msg
      end
    end
  end
end
