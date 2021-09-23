alter table paper change gmt_create
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间';
alter table question change gmt_create
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间';
alter table question_option change gmt_create
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间';
alter table exam.user change gmt_create
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间';
alter table user_pager_record change gmt_create
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间';
alter table user_question_record change gmt_create
    gmt_created datetime default CURRENT_TIMESTAMP not null comment '创建时间';