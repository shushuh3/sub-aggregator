DROP TRIGGER IF EXISTS trg_subscriptions_updated_at ON subscriptions;
DROP TABLE IF EXISTS subscriptions;
DROP FUNCTION IF EXISTS set_updated_at();
DROP EXTENSION IF EXISTS "pgcrypto";
