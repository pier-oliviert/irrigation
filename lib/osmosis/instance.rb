module Osmosis
  class Instance

    def initialize(channel)
      @channel = channel
    end

    def dispatch(cmd)
      @channel.push cmd
    end
  end
end
