## Assignment

1. Faktorial suatu bilangan adalah hasil perkalian bilangan n dengan setiap
bilangan sebelumnya hingga mencapai 1. Faktorial dari 0 adalah 1.
Tugas, buatlah fungsi untuk bilangan bulat dari faktorial n dibagi 2 pangkat n,
dengan sistem pembulatan hasil ke atas.
f(n) = { (n!) ÷ (2
n
) }

2. Dengan menggunakan framework web go fiber, buatlah sistem login dengan
menggunakan data pada redis.
Adapun format redis adalah sebagai berikut:
key: login_<username>
value:
{
“realname”:”Aberto Doni Sianturi”,
“email”:”adss@gmail.com”
“password”:”f7c3bc1d808e0 . . . 441”
}
Note: password = fungsi sha1

# Login Redis App

## Testing

1. Create user: `curl -X POST "http://localhost:3000/create-user?username=testuser"`
2. Login: `curl -X POST http://localhost:3000/login -H "Content-Type: application/json" -d '{"username":"testuser","password":"password123"}'`

👉 [Link Postman Collection](https://web.postman.co/workspace/My-Workspace~5d9abfd2-d6c4-4b08-81b0-6ea048c5fbc9/collection/35986123-7ad20219-fc47-4f1d-8bd5-bb4633a35634?action=share&source=copy-link&creator=35986123)
