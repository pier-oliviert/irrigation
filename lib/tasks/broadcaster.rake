namespace :broadcaster do
  path = Rails.root + 'tmp/pids/broadcaster.pid'
  require 'broadcaster'
  desc 'start broadcaster'
  task :start do
    if File.exists? path
      raise 'Broadcaster is already running.'
    end

#    pid = Process.fork do
      bc = Broadcaster.initialize!
      
#    end

    File.open path, 'w' do |f|
      f << pid
    end
    puts 'Broadcaster is now running.'
  end

  desc 'stop broadcaster'
  task :stop do
    pid = nil
    if File.exists? path
      File.open path, 'r' do |f|
        pid = f.readline.to_i
      end
      File.unlink path
      Process.kill('INT', pid  )
    else
      puts 'Broadcaster is not running.'
    end
  end

end
