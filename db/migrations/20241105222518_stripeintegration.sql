-- Modify "organization_setting_history" table
ALTER TABLE "organization_setting_history" ADD COLUMN "stripe_id" character varying NULL;
-- Modify "organization_settings" table
ALTER TABLE "organization_settings" ADD COLUMN "stripe_id" character varying NULL;
