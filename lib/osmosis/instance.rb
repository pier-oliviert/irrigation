module Osmosis
  class Instance

    def initialize(socket)
      @socket = socket
    end

    def dispatch(cmd)
      return if @socket.nil?
      @socket.send cmd, 0
    end
  end
end
