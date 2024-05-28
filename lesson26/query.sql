
-- 1-tablelarni create qilib oldim
create table student(
    id   uuid primary key default gen_random_uuid() not null,
    name varchar not null,
    age int not null default 18
);

create table course(
    id uuid primary key default gen_random_uuid() not null,
    name varchar not null,
    teacher_name varchar not null
);

create table student_course(
    id uuid primary key default gen_random_uuid() not null,
    student_id uuid references student(id),
    course_id uuid references course(id)
);

create table grade_table(
    id uuid primary key default gen_random_uuid() not null,
    student_course_id uuid references student_course(id),
    grade int not null
);



-- 2-ma'lumotlar insert qilingandan keyin guruhning o'rtacha bahosini har bir guruh bo'yicha chiqarish kodi;

select c.name as course_name,avg(g.grade)
from student as s
join student_course as s1
on s1.student_id=s.id
join course as c
on c.id=s1.course_id
join grade_table as g
on s1.id=g.student_course_id
group by course_name
;

-- 3 eng yosh o'quvchi guruh bo'yicha chiqarilsin chiqarilsin;
-- yani gar bir guruhdan eng yosh o'quvchi(bir nechta bo'lsa har birini)

with new as(select c.name as course_name,min(s.age)
from student as s
join student_course as s1
on s1.student_id=s.id
join course as c
on c.id=s1.course_id
join grade_table as g
on s1.id=g.student_course_id
group by c.name)

select n.course_name,s.name,s.age
from student as s
join student_course as s1
on s1.student_id=s.id
join course as c
on c.id=s1.course_id
join grade_table as g
on s1.id=g.student_course_id
join new as n on n.course_name=c.name
where n.min=s.age;

-- 4 eng yaxshi o'qiydigan guruh chiqarilsin
-- ya'ni, har bir guruhning o'rtacha bahosining eng katta(yaxshi) bo'lgani

select c.name as course_name,round(avg(g.grade)::decimal,2) as score
from student as s
join student_course as s1
on s1.student_id=s.id
join course as c
on c.id=s1.course_id
join grade_table as g
on s1.id=g.student_course_id
group by course_name
order by score desc
limit 1
;

-- 5  guruhdagi eng yaxshi o'qiydigan studentlarni har guruh bo'yicha chiqaring. Agarda
-- eng yaxshi baholar bir nechta kishida bo'lsa, hammasi chiqsin.

with new as(select c.name as course_name,max(g.grade)
from student as s
join student_course as s1
on s1.student_id=s.id
join course as c
on c.id=s1.course_id
join grade_table as g
on s1.id=g.student_course_id
group by c.name)

select n.course_name,s.name,g.grade
from student as s
join student_course as s1
on s1.student_id=s.id
join course as c
on c.id=s1.course_id
join grade_table as g
on s1.id=g.student_course_id
join new as n on n.course_name=c.name
where n.max=g.grade
order by n.course_name
;
