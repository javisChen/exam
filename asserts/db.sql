drop table if exists exam.paper;
create table if not exists exam.paper
(
    id bigint unsigned auto_increment comment 'id'
        primary key,
    title varchar(64) default '' not null comment '试卷标题',
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    gmt_modified datetime default CURRENT_TIMESTAMP not null comment '更新时间',
    is_deleted tinyint(1) default 0 not null comment '逻辑删除 0-未删除 1-已删除'
)
    comment '试卷表';

drop table if exists exam.question;
create table if not exists exam.question
(
    id bigint unsigned auto_increment comment 'id'
        primary key,
    paper_id bigint unsigned null comment '试卷id,关联paper.id',
    title varchar(64) default '' not null comment '试题标题',
    type tinyint(1) default 1 not null comment '试题类型 1-单选 2-多选 3-问答',
    score int default 0 not null comment '试题分数',
    answer varchar(256) default '' not null comment '试题答案，问答题答案为空，单选/多选题存储选项的code，以","分割，例如"1,2,3"',
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    gmt_modified datetime default CURRENT_TIMESTAMP not null comment '更新时间',
    is_deleted tinyint(1) default 0 not null comment '逻辑删除 0-未删除 1-已删除'
)
    comment '试题表';


drop table if exists exam.question_option;
create table if not exists exam.question_option
(
    id bigint unsigned auto_increment comment 'id'
        primary key,
    question_id bigint unsigned null comment '试题id,关联question.id',
    title varchar(64) default '' not null comment '选项标题',
    seq int default 1 not null comment '序列',
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    gmt_modified datetime default CURRENT_TIMESTAMP not null comment '更新时间',
    is_deleted tinyint(1) default 0 not null comment '逻辑删除 0-未删除 1-已删除'
)
    comment '试卷表';

drop table if exists exam.user;
create table if not exists exam.user
(
    id bigint unsigned auto_increment comment 'id'
        primary key,
    username varchar(64) default '' not null comment '用户名',
    phone varchar(64) default '' not null comment '手机',
    password varchar(128) default '' not null comment '密码',
    role tinyint(1) default 1 not null comment '角色 1-教师 2-学生',
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    gmt_modified datetime default CURRENT_TIMESTAMP not null comment '更新时间',
    is_deleted tinyint(1) default 0 not null comment '逻辑删除 0-未删除 1-已删除'
)
    comment '用户表';

drop table if exists exam.user_pager_record;
create table if not exists exam.user_pager_record
(
    id bigint unsigned auto_increment comment 'id'
        primary key,
    paper_id bigint unsigned null comment '试卷id,关联paper.id',
    user_id bigint unsigned null comment '用户id，关联user.id',
    score int default 0 not null comment '获得的分数',
    status tinyint(1) default 0 not null comment '状态 0-未完成 1-已完成',
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    gmt_modified datetime default CURRENT_TIMESTAMP not null comment '更新时间',
    is_deleted tinyint(1) default 0 not null comment '逻辑删除 0-未删除 1-已删除'
)
    comment '用户试卷记录';

drop table if exists exam.user_question_record;
create table if not exists exam.user_question_record
(
    id bigint unsigned auto_increment comment 'id'
        primary key,
    question_id bigint unsigned null comment '试题id,关联question.id',
    paper_id bigint unsigned null comment '试卷id,关联paper.id',
    is_right tinyint(1) default 0 not null comment '是否正确 0-否 1-是',
    user_answer varchar(256) default '' not null comment '试题答案，单选/多选题存储选项的code，以","分割，例如"1,2,3"',
    status tinyint(1) default 0 not null comment '答题状态 0-未完成 1-已完成',
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    gmt_modified datetime default CURRENT_TIMESTAMP not null comment '更新时间',
    is_deleted tinyint(1) default 0 not null comment '逻辑删除 0-未删除 1-已删除'
)
    comment '用户试题答题记录';

