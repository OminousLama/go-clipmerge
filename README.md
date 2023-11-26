
<h1 align="center">
  <br>
  <a href="https://github.com/OminousLama/go-clipmerge"><img src="./docs/res/clipmerge-icon.svg" alt="Clipmerge Icon" width="200"></a>
  <br>
  Clipmerge
  <br>
</h1>

<h4 align="center">Easily merge video clips in the CLI.</h4>

## Getting started

### Prerequisites

- ffmpeg
```bash
# Fedora
sudo dnf install ffmpeg

# Debian
sudo apt install ffmpeg
```


### How To Use

Grab the latest release from the [release page](https://github.com/OminousLama/go-clipmerge/releases/latest).

#### Usage

`clipmerge <output_file> <input_file_1> <input_file_2> <input_file_3> ...`

#### Examples

`clipmerge output.mp4 input1.mp4 input2.mp4 input3.mp4`

This command...
- takes the files `input1.mp4`, `input2.mp4` and `input3.mp4`
- merges them
- and saves the result in `output.mp4`

## Credits

This software uses the following open source packages:

- [Best README Template](https://github.com/othneildrew/Best-README-Template)
- [Go](https://github.com/golang)
- [Remix Icon](https://remixicon.com/)
- [ffmpeg](https://ffmpeg.org/)


## License

GPL-3.0 license
