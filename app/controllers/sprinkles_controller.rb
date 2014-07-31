class SprinklesController < ApplicationController
  def create
    @sprinkle = zone.sprinkles.create! sprinkle_params
    cmd = {
      action: {
        name: "open",
        id: @sprinkle.zone.gpio
      }
    }
    osmosis.dispatch JSON.generate(cmd)

    respond_to do |format|
      format.html
      format.js do
        render nothing: true
      end
    end
  end

  def destroy
    @sprinkle = Sprinkle.find(params[:id])
    @sprinkle.ends_at = Time.now
    @sprinkle.save!
    cmd = {
      action: {
        name: "close",
        id: @sprinkle.zone.gpio
      }
    }
    osmosis.dispatch JSON.generate(cmd)
    respond_to do |format|
      format.html
      format.js do
        render nothing: true
      end
    end
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
