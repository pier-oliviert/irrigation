require 'broadcaster/osmosis'

require_relative '../../app/models/sprinkle'
require_relative '../../app/models/zone'

module Broadcaster
  class Manager
    def initialize
      @connections = []
      @osmosis = Broadcaster::Osmosis.new
      @osmosis.run! do |cmd|
        sprinkle = Sprinkle.find(cmd.last)
        @connections.each do |c|
          c.send "zone:#{cmd.first}:#{sprinkle.zone.id}"
        end
      end
    end

    def register(conn)
      @connections.push conn
      conn.onclose do
        @connections.delete(conn)
      end
    end

    def configure!
      ActiveRecord::Base.configurations = Rails.application.config.database_configuration
      ActiveRecord::Base.establish_connection
    end
  end
end
