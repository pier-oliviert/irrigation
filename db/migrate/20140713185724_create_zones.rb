class CreateZones < ActiveRecord::Migration
  def change
    create_table :zones do |t|
      t.integer :gpio, null: false
      t.string :name, null: false
      t.boolean :opened, default: false
      t.timestamps
    end
  end
end
