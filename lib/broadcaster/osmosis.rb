module Broadcaster
  class Osmosis
    attr_reader :error

    def initialize
    end

    def run!
      raise 'Need a block to run Osmosis' unless block_given?
      Thread.new do
        loop do
          msg = parse(socket.recv(1024 * 32))
          unless msg.nil?
            yield(msg)
          end
        end
      end
    end

    def parse(msg)
      return if msg.empty?
      puts msg
      msg.split(":")
    end

    def socket
      @socket ||= UNIXSocket.new(path)
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
