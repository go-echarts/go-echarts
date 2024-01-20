# Dive into go-echarts

# Summary

?> go-echarts is a golang lib providing the power to create [echarts](https://echarts.apache.org/) styling charts in go.

---

# What go-echarts does:

- Provide the generic HTML templates with echarts resources.
- Generate echarts required options and datasets based on go programs.
- Mount options to echarts' instance and render all in one place.

# Where to find options I need?

go-echarts' options are almost align to [echarts' official options](https://echarts.apache.org/en/option.html#title).  
If you want to find the options of your specific charts, all in the echarts options list.  
And if you find some options go-echarts hasn't support yet, kindly create a PR or raise an option request.

# Where to find custom echarts resources

Currently, we support echarts' version is `v5.4.3` and
the compatible version `v4.3` for those 3d charts and customized charts, such as word-cloud.  
All the resources are hosted in [go-echarts-assets](https://github.com/go-echarts/go-echarts-assets) by go-echarts
itself.

If you want to use your custom echarts resources.
You could find in the [jsdelivr of echarts](https://www.jsdelivr.com/package/npm/echarts) or other CDN providers.
Besides, you can set your own assets host as well.

We support 2 ways to customize the echarts resources.

- Only change the `Asserts Host`, which needs the echarts resources naming `echarts.min.js`
- Clean all the preset assets and add custom resources all by yourself.



