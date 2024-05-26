-- many to many kop bo'lishi mumkin bo'lgan narsalarga ishlatadi misol uchun bizda mahsulot bor 
-- va bizni mahsulotimizni ko'p odamda bor va bitta odamni ko'p boshqa mahsulotlari bor bu many to manyga misol

create table user(
    id    uuid primary key not null default gen_random_uuid(),
    name  varchar   not null,
);


create table car
(
    id     uuid primary key not null default gen_random_uuid(),
    name   varchar          not null,
    year    int not null,
    color varchar not null default 'white',
);

create table saver(
    id serial primary key,
    car_id uuid not null,
    user_id uuid not null,
    foreign key(car_id) references car(id),
    foreign key(user_id) references user(id)
);
