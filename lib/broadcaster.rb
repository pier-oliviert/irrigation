require 'broadcaster/manager'
module Broadcaster
  def self.initialize!
    @manager = Broadcaster::Manager.new
    @manager.configure!
    EM.run do
      listen!
    end
  end

  def self.listen!
    Thread.new do
      EM::WebSocket.run(:host => "0.0.0.0", :port => 21343) do |ws|
        @manager.register(ws)
      end
    end
  end
end
