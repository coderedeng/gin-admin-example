/*
 Navicat Premium Data Transfer

 Source Server         : PgSQL
 Source Server Type    : PostgreSQL
 Source Server Version : 160001
 Source Host           : localhost:5432
 Source Catalog        : gpa
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160001
 File Encoding         : 65001

 Date: 06/03/2024 21:28:08
*/


-- ----------------------------
-- Sequence structure for apis_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "apis_id_seq";
CREATE SEQUENCE "apis_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for authorities_authority_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "authorities_authority_id_seq";
CREATE SEQUENCE "authorities_authority_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for base_menu_btn_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "base_menu_btn_id_seq";
CREATE SEQUENCE "base_menu_btn_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for base_menus_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "base_menus_id_seq";
CREATE SEQUENCE "base_menus_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for base_menus_parameter_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "base_menus_parameter_id_seq";
CREATE SEQUENCE "base_menus_parameter_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for casbin_rule_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "casbin_rule_id_seq";
CREATE SEQUENCE "casbin_rule_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for jwt_blacklists_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "jwt_blacklists_id_seq";
CREATE SEQUENCE "jwt_blacklists_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "users_id_seq";
CREATE SEQUENCE "users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for apis
-- ----------------------------
DROP TABLE IF EXISTS "apis";
CREATE TABLE "apis" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('apis_id_seq'::regclass),
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz",
  "path" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "description" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "api_group" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "method" "pg_catalog"."text" COLLATE "pg_catalog"."default" DEFAULT 'POST'::text
)
;
COMMENT ON COLUMN "apis"."path" IS 'api路径';
COMMENT ON COLUMN "apis"."description" IS 'api中文描述';
COMMENT ON COLUMN "apis"."api_group" IS 'api组';
COMMENT ON COLUMN "apis"."method" IS '方法';

-- ----------------------------
-- Records of apis
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for authorities
-- ----------------------------
DROP TABLE IF EXISTS "authorities";
CREATE TABLE "authorities" (
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz",
  "authority_id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('authorities_authority_id_seq'::regclass),
  "authority_name" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "parent_id" "pg_catalog"."int8",
  "default_router" "pg_catalog"."text" COLLATE "pg_catalog"."default" DEFAULT 'dashboard'::text
)
;
COMMENT ON COLUMN "authorities"."authority_id" IS '角色ID';
COMMENT ON COLUMN "authorities"."authority_name" IS '角色名';
COMMENT ON COLUMN "authorities"."parent_id" IS '父角色ID';
COMMENT ON COLUMN "authorities"."default_router" IS '默认菜单';

-- ----------------------------
-- Records of authorities
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for authority_btn
-- ----------------------------
DROP TABLE IF EXISTS "authority_btn";
CREATE TABLE "authority_btn" (
  "authority_id" "pg_catalog"."int8",
  "sys_menu_id" "pg_catalog"."int8",
  "sys_base_menu_btn_id" "pg_catalog"."int8"
)
;
COMMENT ON COLUMN "authority_btn"."authority_id" IS '角色ID';
COMMENT ON COLUMN "authority_btn"."sys_menu_id" IS '菜单ID';
COMMENT ON COLUMN "authority_btn"."sys_base_menu_btn_id" IS '菜单按钮ID';

-- ----------------------------
-- Records of authority_btn
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for authority_menus
-- ----------------------------
DROP TABLE IF EXISTS "authority_menus";
CREATE TABLE "authority_menus" (
  "sys_base_menu_id" "pg_catalog"."int8" NOT NULL,
  "sys_authority_authority_id" "pg_catalog"."int8" NOT NULL
)
;
COMMENT ON COLUMN "authority_menus"."sys_authority_authority_id" IS '角色ID';

-- ----------------------------
-- Records of authority_menus
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for base_menu_btn
-- ----------------------------
DROP TABLE IF EXISTS "base_menu_btn";
CREATE TABLE "base_menu_btn" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('base_menu_btn_id_seq'::regclass),
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz",
  "name" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "desc" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "sys_base_menu_id" "pg_catalog"."int8"
)
;
COMMENT ON COLUMN "base_menu_btn"."name" IS '按钮关键key';
COMMENT ON COLUMN "base_menu_btn"."sys_base_menu_id" IS '菜单ID';

-- ----------------------------
-- Records of base_menu_btn
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for base_menus
-- ----------------------------
DROP TABLE IF EXISTS "base_menus";
CREATE TABLE "base_menus" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('base_menus_id_seq'::regclass),
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz",
  "menu_level" "pg_catalog"."int8",
  "parent_id" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "path" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "name" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "hidden" "pg_catalog"."bool",
  "component" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "sort" "pg_catalog"."int8",
  "active_name" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "keep_alive" "pg_catalog"."bool",
  "default_menu" "pg_catalog"."bool",
  "title" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "icon" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "close_tab" "pg_catalog"."bool"
)
;
COMMENT ON COLUMN "base_menus"."parent_id" IS '父菜单ID';
COMMENT ON COLUMN "base_menus"."path" IS '路由path';
COMMENT ON COLUMN "base_menus"."name" IS '路由name';
COMMENT ON COLUMN "base_menus"."hidden" IS '是否在列表隐藏';
COMMENT ON COLUMN "base_menus"."component" IS '对应前端文件路径';
COMMENT ON COLUMN "base_menus"."sort" IS '排序标记';
COMMENT ON COLUMN "base_menus"."active_name" IS '高亮菜单';
COMMENT ON COLUMN "base_menus"."keep_alive" IS '是否缓存';
COMMENT ON COLUMN "base_menus"."default_menu" IS '是否是基础路由（开发中）';
COMMENT ON COLUMN "base_menus"."title" IS '菜单名';
COMMENT ON COLUMN "base_menus"."icon" IS '菜单图标';
COMMENT ON COLUMN "base_menus"."close_tab" IS '自动关闭tab';

-- ----------------------------
-- Records of base_menus
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for base_menus_parameter
-- ----------------------------
DROP TABLE IF EXISTS "base_menus_parameter";
CREATE TABLE "base_menus_parameter" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('base_menus_parameter_id_seq'::regclass),
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz",
  "sys_base_menu_id" "pg_catalog"."int8",
  "type" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "key" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "value" "pg_catalog"."text" COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "base_menus_parameter"."type" IS '地址栏携带参数为params还是query';
COMMENT ON COLUMN "base_menus_parameter"."key" IS '地址栏携带参数的key';
COMMENT ON COLUMN "base_menus_parameter"."value" IS '地址栏携带参数的值';

-- ----------------------------
-- Records of base_menus_parameter
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS "casbin_rule";
CREATE TABLE "casbin_rule" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('casbin_rule_id_seq'::regclass),
  "ptype" "pg_catalog"."varchar" COLLATE "pg_catalog"."default",
  "v0" "pg_catalog"."varchar" COLLATE "pg_catalog"."default",
  "v1" "pg_catalog"."varchar" COLLATE "pg_catalog"."default",
  "v2" "pg_catalog"."varchar" COLLATE "pg_catalog"."default",
  "v3" "pg_catalog"."varchar" COLLATE "pg_catalog"."default",
  "v4" "pg_catalog"."varchar" COLLATE "pg_catalog"."default",
  "v5" "pg_catalog"."varchar" COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for data_authority_ids
-- ----------------------------
DROP TABLE IF EXISTS "data_authority_ids";
CREATE TABLE "data_authority_ids" (
  "sys_authority_authority_id" "pg_catalog"."int8" NOT NULL,
  "data_authority_id_authority_id" "pg_catalog"."int8" NOT NULL
)
;
COMMENT ON COLUMN "data_authority_ids"."sys_authority_authority_id" IS '角色ID';
COMMENT ON COLUMN "data_authority_ids"."data_authority_id_authority_id" IS '角色ID';

-- ----------------------------
-- Records of data_authority_ids
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS "jwt_blacklists";
CREATE TABLE "jwt_blacklists" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('jwt_blacklists_id_seq'::regclass),
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz",
  "jwt" "pg_catalog"."text" COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "jwt_blacklists"."jwt" IS 'jwt';

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_authorities
-- ----------------------------
DROP TABLE IF EXISTS "user_authorities";
CREATE TABLE "user_authorities" (
  "sys_user_id" "pg_catalog"."int8" NOT NULL,
  "sys_authority_authority_id" "pg_catalog"."int8" NOT NULL
)
;
COMMENT ON COLUMN "user_authorities"."sys_authority_authority_id" IS '角色ID';

-- ----------------------------
-- Records of user_authorities
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "id" "pg_catalog"."int8" NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "created_at" "pg_catalog"."timestamptz",
  "updated_at" "pg_catalog"."timestamptz",
  "deleted_at" "pg_catalog"."timestamptz",
  "uuid" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "username" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "password" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "nick_name" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "side_mode" "pg_catalog"."text" COLLATE "pg_catalog"."default" DEFAULT 'dark'::text,
  "header_img" "pg_catalog"."text" COLLATE "pg_catalog"."default" DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg'::text,
  "base_color" "pg_catalog"."text" COLLATE "pg_catalog"."default" DEFAULT '#fff'::text,
  "active_color" "pg_catalog"."text" COLLATE "pg_catalog"."default" DEFAULT '#1890ff'::text,
  "authority_id" "pg_catalog"."int8" DEFAULT 888,
  "phone" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "email" "pg_catalog"."text" COLLATE "pg_catalog"."default",
  "enable" "pg_catalog"."int8" DEFAULT 1
)
;
COMMENT ON COLUMN "users"."uuid" IS '用户UUID';
COMMENT ON COLUMN "users"."username" IS '用户登录名';
COMMENT ON COLUMN "users"."password" IS '用户登录密码';
COMMENT ON COLUMN "users"."nick_name" IS '用户昵称';
COMMENT ON COLUMN "users"."side_mode" IS '用户侧边主题';
COMMENT ON COLUMN "users"."header_img" IS '用户头像';
COMMENT ON COLUMN "users"."base_color" IS '基础颜色';
COMMENT ON COLUMN "users"."active_color" IS '活跃颜色';
COMMENT ON COLUMN "users"."authority_id" IS '用户角色ID';
COMMENT ON COLUMN "users"."phone" IS '用户手机号';
COMMENT ON COLUMN "users"."email" IS '用户邮箱';
COMMENT ON COLUMN "users"."enable" IS '用户是否被冻结 1正常 2冻结';

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "apis_id_seq"
OWNED BY "apis"."id";
SELECT setval('"apis_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "authorities_authority_id_seq"
OWNED BY "authorities"."authority_id";
SELECT setval('"authorities_authority_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "base_menu_btn_id_seq"
OWNED BY "base_menu_btn"."id";
SELECT setval('"base_menu_btn_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "base_menus_id_seq"
OWNED BY "base_menus"."id";
SELECT setval('"base_menus_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "base_menus_parameter_id_seq"
OWNED BY "base_menus_parameter"."id";
SELECT setval('"base_menus_parameter_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "casbin_rule_id_seq"
OWNED BY "casbin_rule"."id";
SELECT setval('"casbin_rule_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "jwt_blacklists_id_seq"
OWNED BY "jwt_blacklists"."id";
SELECT setval('"jwt_blacklists_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "users_id_seq"
OWNED BY "users"."id";
SELECT setval('"users_id_seq"', 1, false);

-- ----------------------------
-- Indexes structure for table apis
-- ----------------------------
CREATE INDEX "idx_apis_deleted_at" ON "apis" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table apis
-- ----------------------------
ALTER TABLE "apis" ADD CONSTRAINT "apis_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table authorities
-- ----------------------------
ALTER TABLE "authorities" ADD CONSTRAINT "authorities_pkey" PRIMARY KEY ("authority_id");

-- ----------------------------
-- Primary Key structure for table authority_menus
-- ----------------------------
ALTER TABLE "authority_menus" ADD CONSTRAINT "authority_menus_pkey" PRIMARY KEY ("sys_base_menu_id", "sys_authority_authority_id");

-- ----------------------------
-- Indexes structure for table base_menu_btn
-- ----------------------------
CREATE INDEX "idx_base_menu_btn_deleted_at" ON "base_menu_btn" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table base_menu_btn
-- ----------------------------
ALTER TABLE "base_menu_btn" ADD CONSTRAINT "base_menu_btn_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table base_menus
-- ----------------------------
CREATE INDEX "idx_base_menus_deleted_at" ON "base_menus" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table base_menus
-- ----------------------------
ALTER TABLE "base_menus" ADD CONSTRAINT "base_menus_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table base_menus_parameter
-- ----------------------------
CREATE INDEX "idx_base_menus_parameter_deleted_at" ON "base_menus_parameter" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table base_menus_parameter
-- ----------------------------
ALTER TABLE "base_menus_parameter" ADD CONSTRAINT "base_menus_parameter_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table casbin_rule
-- ----------------------------
CREATE UNIQUE INDEX "idx_casbin_rule" ON "casbin_rule" USING btree (
  "ptype" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v0" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v1" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v2" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v3" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v4" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v5" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table casbin_rule
-- ----------------------------
ALTER TABLE "casbin_rule" ADD CONSTRAINT "casbin_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table data_authority_ids
-- ----------------------------
ALTER TABLE "data_authority_ids" ADD CONSTRAINT "data_authority_ids_pkey" PRIMARY KEY ("sys_authority_authority_id", "data_authority_id_authority_id");

-- ----------------------------
-- Indexes structure for table jwt_blacklists
-- ----------------------------
CREATE INDEX "idx_jwt_blacklists_deleted_at" ON "jwt_blacklists" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table jwt_blacklists
-- ----------------------------
ALTER TABLE "jwt_blacklists" ADD CONSTRAINT "jwt_blacklists_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table user_authorities
-- ----------------------------
ALTER TABLE "user_authorities" ADD CONSTRAINT "user_authorities_pkey" PRIMARY KEY ("sys_user_id", "sys_authority_authority_id");

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE INDEX "idx_users_deleted_at" ON "users" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_users_username" ON "users" USING btree (
  "username" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_users_uuid" ON "users" USING btree (
  "uuid" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "users" ADD CONSTRAINT "users_authority_id_key" UNIQUE ("authority_id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table authorities
-- ----------------------------
ALTER TABLE "authorities" ADD CONSTRAINT "fk_users_authority" FOREIGN KEY ("authority_id") REFERENCES "users" ("authority_id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table authority_btn
-- ----------------------------
ALTER TABLE "authority_btn" ADD CONSTRAINT "fk_authority_btn_sys_base_menu_btn" FOREIGN KEY ("sys_base_menu_btn_id") REFERENCES "base_menu_btn" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table authority_menus
-- ----------------------------
ALTER TABLE "authority_menus" ADD CONSTRAINT "fk_authority_menus_sys_authority" FOREIGN KEY ("sys_authority_authority_id") REFERENCES "authorities" ("authority_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "authority_menus" ADD CONSTRAINT "fk_authority_menus_sys_base_menu" FOREIGN KEY ("sys_base_menu_id") REFERENCES "base_menus" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table base_menu_btn
-- ----------------------------
ALTER TABLE "base_menu_btn" ADD CONSTRAINT "fk_base_menus_menu_btn" FOREIGN KEY ("sys_base_menu_id") REFERENCES "base_menus" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table base_menus_parameter
-- ----------------------------
ALTER TABLE "base_menus_parameter" ADD CONSTRAINT "fk_base_menus_parameters" FOREIGN KEY ("sys_base_menu_id") REFERENCES "base_menus" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table data_authority_ids
-- ----------------------------
ALTER TABLE "data_authority_ids" ADD CONSTRAINT "fk_data_authority_ids_data_authority_id" FOREIGN KEY ("data_authority_id_authority_id") REFERENCES "authorities" ("authority_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "data_authority_ids" ADD CONSTRAINT "fk_data_authority_ids_sys_authority" FOREIGN KEY ("sys_authority_authority_id") REFERENCES "authorities" ("authority_id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table user_authorities
-- ----------------------------
ALTER TABLE "user_authorities" ADD CONSTRAINT "fk_user_authorities_sys_authority" FOREIGN KEY ("sys_authority_authority_id") REFERENCES "authorities" ("authority_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "user_authorities" ADD CONSTRAINT "fk_user_authorities_sys_user" FOREIGN KEY ("sys_user_id") REFERENCES "users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
