## Lệnh git cơ bản 

1. Tạo branch mới từ master branch (luôn luôn tạo nhánh con từ master branch)
git checkout -b {branch_name}
git checkout -b add-template-go

2. Lệnh kiểm tra branch hiện tại 
git branch

3. Add File lên github
 git add .

4. Dùng để tạo lệnh ghi chú lên github
git commit -m "Them template learning go" 

5. Đẩy code lên github 
  git push --set-upstream origin add-template-go

6. Mở github lên https://github.com/oringv/goapp-giahuy
- Click vào Compare pull request
- Tạo mới pull requrest (PR)
- Tạo tiêu đề của PR: vidu Add template learning go
- Tạo description của PR
- Click vào nút tạo PR
- Xem commit đã có chưa 
- Chọn reviewer cho PR

## Sau khi tạo pull requrest (PR1) thành công, tiếp tục commit code mới 

- Ví dụ PR#1: https://github.com/oringv/goapp-giahuy/pull/1
- Thêm thư mục day3, và cần commit code vừa thay đổi cần thực hiện
- Lập lại bước số 3 
git add .
git commit -m "Add day3 folder"
git push 

## Tạo branch mới để học go 

- iểm tra branch hiện tại 
git branch
- Phải chuyển về branch master trước khi tạo branch con 
git checkout master  
- Kiểm tra branch hiện tại 
git branch 
-> Branch hiện tại đang là master
- Lấy code mới nhất từ github về branch master 
 git pull
- Tạo branch mới để học video từ số 1-10
git checkout -b video-day1
- Kiểm tra branch hiện tại 
