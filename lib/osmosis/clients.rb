module Osmosis
  class Clients
    attr_reader :channel

    def initialize
      @channel = EM::Channel.new
    end

    def start!
      socket do |ws|
        sid = nil
        ws.onopen do
          sid = @channel.subscribe do |msg|
            ws.send(msg)
          end
        end

        ws.onerror do |error|
          puts error
        end

        ws.onclose do
          unless sid.nil?
            @channel.unsubscribe sid
          end
        end
      end
    end

    protected

    def socket(&block)
      @socket ||= begin
        raise 'Need a block to start the socket' unless block_given?
        Thread.new do
          EM::WebSocket.start(host: "0.0.0.0", port: 21343) do |ws|
            block.call(ws)
          end
        end
      end
    end
  end
end
