#! /usr/bin/env ruby

require "socket"
require "json"

@clients = []
@pins = [
  {state: 0, id:3},
  {state: 0, id:5},
  {state: 0, id:8},
  {state: 0, id:10},
  {state: 0, id:11},
  {state: 0, id:12},
  {state: 0, id:13},
  {state: 0, id:15},
  {state: 0, id:16},
  {state: 0, id:18},
  {state: 0, id:19},
  {state: 0, id:21},
  {state: 0, id:22},
  {state: 0, id:23},
  {state: 0, id:24},
  {state: 0, id:26},
]

serv = UNIXServer.new("/tmp/gobble.sock")

def list_pins(pins)
  msg = JSON.generate(pins)
  puts msg
  @clients.each do |c|
    c.send msg, 0
  end
end

begin
while true
  c = serv.accept
  @clients << c
  Thread.new do
    running = true
    while running do
      begin
        msg, sender, flags, _ = c.recvmsg
        data = JSON.parse(msg)
        action = data['action']

        unless action.nil?
          case action['name']
          when 'open'
            puts "Opening pin ##{action['id']}"
            @pins.each do |p|
              if p[:id] == action['id']
                p[:state] = 1
                puts p
              end
            end
            list_pins(@pins)
          when 'close'
            puts "Closing pin ##{action['id']}"
            @pins.each do |p|
              if p[:id] == action['id']
                p[:state] = 0
              end
            end
            list_pins(@pins)
          else
            list_pins(@pins)
          end
        end
      rescue StandardError => e
        running false
      end
    end
    @clients.delete(c)
    c.close()
  end
end

rescue SystemExit, Interrupt
  File.unlink serv.path
  raise
end
