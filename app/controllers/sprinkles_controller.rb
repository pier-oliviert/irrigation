class SprinklesController < ApplicationController
  def create
    @sprinkle = zone.sprinkles.create! sprinkle_params
    osmosis.dispatch "open:#{@sprinkle.id}"
  end

  def destroy
    @sprinkle = Sprinkle.find(params[:id])
    @sprinkle.ends_at = Time.now
    @sprinkle.save!
    osmosis.dispatch "close:#{@sprinkle.id}"
  end

  protected

  def sprinkle_params
    params.require(:sprinkle).permit(:duration)
  end

  def zone
    @zone ||= begin
      Zone.find(params[:zone_id])
    end
  end
end
