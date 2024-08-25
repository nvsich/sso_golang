insert into apps(id, name, secret)
values (1, 'test name', 'test secret')
on conflict do nothing;