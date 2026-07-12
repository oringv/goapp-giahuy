## lệnh git cơ bản 

1. tạo branch
git checkout -b {branch_name}
git checkout -b add-template-go

2. Lệnh kiểm tra branch hiện tại 
git branch

3.Add File lên github
 git add .

4. Dùng để tạo lệnh ghi chú lên github
git commit -m "Them template learning go" 

5. Đẩy code lên github 
  git push --set-upstream origin add-template-go


## Sau khi tạo pull requrest (PR1) thành công, tiếp tục commit code mới 

- Ví dụ PR#1: https://github.com/oringv/goapp-giahuy/pull/1
- Thêm thư mục day3, và cần commit code vừa thay đổi cần thực hiện

- Lập lại bước số 3 
git add .
git commit -m "Add day3 folder"
git push 


