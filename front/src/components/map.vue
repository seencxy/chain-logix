<template>
    <div id="container"></div>
</template>
 
<script>
import AMapLoader from "@amap/amap-jsapi-loader";
export default {
    name: "map-view",

    mounted() {
        this.initAMap();
    },

    unmounted() {
        this.map?.destroy();
    },
    props: {
        centerCoordinates: Array,
        start: Array,
        end: Array,
        // 可以添加更多的props来接收其他数据
    },

    methods: {
        initAMap() {
            AMapLoader.load({
                key: "25125b66d60b8e29dfdde781c72bf6cd", // 替换为您的高德地图Web端开发者Key
                version: "2.0",
                // 下面是插件列表
                plugins: [
                    'AMap.Autocomplete', 'AMap.PlaceSearch', 'AMap.Scale',
                    'AMap.PolyEditor', 'AMap.CircleEditor', 'AMap.Geocoder'
                ],
            })
                .then((AMap) => {
                    this.map = new AMap.Map("container", {
                        viewMode: "3D",
                        resizeEnable: true,
                        zoom: 17,
                        zooms: [2, 20],
                        rotation: -15,
                        pitch: 50,
                        rotateEnable: true,
                        pitchEnable: true,
                        center: this.centerCoordinates,
                    });

                    // 创建 geocoder 实例
                    this.geocoder = new AMap.Geocoder({
                        city: "全国", // 或其他默认值
                    });


                    // 在这里添加额外的地图配置和功能
                    // 例如添加标记、事件监听器等
                    this.addCustomFeatures(AMap, this.map);
                })
                .catch((e) => {
                    console.log(e);
                });
        },
        addCustomFeatures(AMap, map) {
            // 以 icon URL 的形式创建一个途经点
            var viaMarker = new AMap.Marker({
                position: new AMap.LngLat(this.centerCoordinates[0], this.centerCoordinates[1]),
                // 将一张图片的地址设置为 icon
                icon: '//a.amap.com/jsapi_demos/static/demo-center/icons/dir-via-marker.png',
                // 设置了 icon 以后，设置 icon 的偏移量，以 icon 的 [center bottom] 为原点
                offset: new AMap.Pixel(-13, -30)
            });

            // 创建一个 Icon
            var startIcon = new AMap.Icon({
                // 图标尺寸
                size: new AMap.Size(25, 34),
                // 图标的取图地址
                image: '//a.amap.com/jsapi_demos/static/demo-center/icons/dir-marker.png',
                // 图标所用图片大小
                imageSize: new AMap.Size(135, 40),
                // 图标取图偏移量
                imageOffset: new AMap.Pixel(-9, -3)
            });

            // 将 icon 传入 marker
            var startMarker = new AMap.Marker({
                position: new AMap.LngLat(this.start[0], this.start[1]),
                icon: startIcon,
                offset: new AMap.Pixel(-13, -30)
            });

            // 创建一个 icon
            var endIcon = new AMap.Icon({
                size: new AMap.Size(25, 34),
                image: '//a.amap.com/jsapi_demos/static/demo-center/icons/dir-marker.png',
                imageSize: new AMap.Size(135, 40),
                imageOffset: new AMap.Pixel(-95, -3)
            });

            // 将 icon 传入 marker
            var endMarker = new AMap.Marker({
                position: new AMap.LngLat(this.end[0], this.end[1]),
                icon: endIcon,
                offset: new AMap.Pixel(-13, -30)
            });

            // 将 markers 添加到地图
            map.add([viaMarker, startMarker, endMarker]);
            // 在添加标记后，将地图中心设置为当前位置
            map.setCenter(new AMap.LngLat(this.centerCoordinates[0], this.centerCoordinates[1]));
        },
        updateMap() {
            if (this.map) {
                // 在这里执行地图更新的相关操作
                // 例如清除之前的标记，然后重新添加标记
                this.map.clearMap();

                // 添加新的标记
                this.addCustomFeatures(AMap, this.map);
            }
        },
    },
    watch: {
        centerCoordinates: function (newCoordinates) {
            this.updateMap();
        },
        start: function (newStart) {
            this.updateMap();
        },
        end: function (newEnd) {
            this.updateMap();
        }
    },
};
</script>
  
<style scoped>
#container {
    width: 100%;
    height: 100%;
}
</style>
  