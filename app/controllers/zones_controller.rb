class ZonesController < ApplicationController
  def index
    @zones = Zone.all
    render 'onboarding' and return if @zones.count == 0
  end

  def create
    @zone = Zone.create(zone_params)
    redirect_to :root
  end

  protected

  def zone_params
    params.require(:zone).permit(:name, :gpio)
  end
end
