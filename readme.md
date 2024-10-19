# golang-uniswapV3

## 功能

1. 使用uniswapV3 获取报价

```sh
go run cmd/quoter/main.go
```

2. 使用uniswapV3 实现任意两个加密货币的兑换

**项目根目录下创建 .env 文件**

``` txt
PRIVATE_KEY=your_private_key
INFURA_PROJECT_ID=your_infura_project_id
```

```sh
go run cmd/swap/main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.