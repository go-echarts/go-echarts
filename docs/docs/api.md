---
id: api
title: API 设计
sidebar_label: API 设计
---

go-echarts 秉承着 API 设计要简洁的原则，对所有图表只提供了下面的接口


### Render(w ...io.Writer)
负责渲染图表，支持传入多个实现了 io.Writer 接口的对象

Add()

AddXAxis()

AddYAxis()

Overlap()

// 设置全局配置项
func SetGlobalOptions(options ...globalOptser) *RectChart {}
// 设置 Series 配置项
func SetSeriesOptions(options ...seriesOptser) {}
// 渲染图表
func Render(w ...io.Writer) error {}