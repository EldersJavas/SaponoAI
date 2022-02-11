# 图片

## 命令
```sh
realesrgan-ncnn-vulkan -i input.jpg -o output.png -n model_name
```

# 视频

## 命令
```sh
ffmpeg -i onepiece_demo.mp4 -qscale:v 1 -qmin 1 -qmax 1 -vsync 0 tmp_frames/frame%08d.png //视频转图片
realesrgan-ncnn-vulkan -i tmp_frames -o out_frames -n RealESRGANv2-animevideo-xsx2 -s 2 -f jpg //AI处理图片
ffmpeg -i onepiece_demo.mp4 //查看原视频码率
ffmpeg -r [码率] -i out_frames/frame%08d.jpg -i onepiece_demo.mp4 -map 0:v:0 -map 1:a:0 -c:a copy -c:v libx264 -r [码率] -pix_fmt yuv420p output_w_audio.mp4 //合并视频并拷贝声音
```



# help
```bash
Usage: realesrgan-ncnn-vulkan.exe -i infile -o outfile [options]...

  -h                   show this help
  -v                   verbose output
  -i input-path        input image path (jpg/png/webp) or directory
  -o output-path       output image path (jpg/png/webp) or directory
  -s scale             upscale ratio (4, default=4)
  -t tile-size         tile size (>=32/0=auto, default=0) can be 0,0,0 for multi-gpu
  -m model-path        folder path to pre-trained models(default=models)
  -n model-name        model name (default=realesrgan-x4plus, can be realesrgan-x4plus | realesrgan-x4plus-anime | realesrnet-x4plus)
  -g gpu-id            gpu device to use (default=0) can be 0,1,2 for multi-gpu
  -j load:proc:save    thread count for load/proc/save (default=1:2:2) can be 1:2,2,2:2 for multi-gpu
  -x                   enable tta mode
  -f format            output image format (jpg/png/webp, default=ext/png)

```

