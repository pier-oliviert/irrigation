module Osmosis
  class Backend
    attr_reader :channel

    def initialize
      @channel = EM::Channel.new
      @channel.subscribe do |msg|
        socket.send msg, 0
      end
    end

    def start!
      raise 'Need a block to start the socket' unless block_given?
      Thread.new do
        loop do
          d, sender, flags, _ = socket.recvmsg
          yield(d)
        end
      end
    end

    def socket
      @socket ||= UNIXSocket.new(path)
    end

    def connected?
      error.nil?
    end

    def path
      Rails.application.paths['sockets/osmosis'].existent.first
    end

  end
end
