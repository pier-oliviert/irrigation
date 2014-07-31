class Zone < ActiveRecord::Base
  has_many :sprinkles

  def closing_at
    active_sprinkles = sprinkles.active
    return unless active_sprinkles.any?
    will_close_at = nil
    active_sprinkles.each do |s|
      if will_close_at.nil? || s.ends_at > will_close_at
        will_close_at = s.ends_at
      end
    end

    will_close_at
  end
end
