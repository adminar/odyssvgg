# odyssey
GO + VUE3编写的运维平台

后端:使用GOGIN框架

VUE3:使用Vue3+TypeScript+Vite

```
##后端调试
go mod tidy
GO_ENV=dev go run cmd/server/main.go
##启动后端后尝试
curl 127.0.0.1:8000/man
##前端调试
pnpm install
cd web && npm run dev
```

预期实现功能CI/CD、资产管理、工单系统

目前已有版本写在公司私有仓库，没有脱敏，正在慢慢迁移到github

代码初学者，有什么代码上的意见可以直接提出来，听劝

觉得有点东西的希望家人们可以给个🌟
