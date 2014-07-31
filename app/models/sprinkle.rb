class Sprinkle < ActiveRecord::Base
  belongs_to :zone

  scope :active, -> { where("sprinkles.ends_at > ?", Time.now)}

  def duration=(seconds)
    self.ends_at = seconds.to_i.seconds.from_now
  end

  def duration
    if ends_at.nil?
      30
    else
      return ends_at - Time.now
    end
  end

  def remaining
    ends_at - Time.now
  end
end
