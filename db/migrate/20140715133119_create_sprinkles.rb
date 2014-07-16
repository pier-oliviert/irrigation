class CreateSprinkles < ActiveRecord::Migration
  def change
    create_table :sprinkles do |t|
      t.datetime :ends_at, null: false
      t.references :zone, null: false
      t.timestamps
    end
  end
end
