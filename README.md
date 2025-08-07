# MuxFarm

A distributed media processing system for transmuxing, transcoding, and packaging media files from various sources into streaming-ready formats.

## Overview

MuxFarm provides a scalable architecture for processing media content, supporting ingestion from multiple storage backends and conversion to modern streaming formats like MPEG-DASH and HLS. Built with Go and leveraging FFmpeg for media operations, it's designed to handle large-scale media workflows.

## Key Features

### Media Processing

- **Transmuxing**: Container format conversion (MP4, WebM, FLV, MPEG)
- **Transcoding**: Video (H.264, VP9, AV1) and audio (AAC, MP3, Opus) codec support
- **Adaptive Streaming**: MPEG-DASH and HLS packaging with segmentation
- **Quality Control**: Bitrate and resolution management

### FFmpeg Integration

MuxFarm leverages FFmpeg's powerful multimedia framework through a comprehensive Go wrapper:

- **ffprobe**: Media analysis and metadata extraction
  - Container format detection
  - Stream information parsing
  - Duration and bitrate analysis
  - Codec identification

- **ffmpeg**: Core media processing engine
  - Stream mapping and filtering
  - Real-time transcoding with hardware acceleration support
  - Custom filter graphs for complex transformations
  - Pipe-based I/O for memory-efficient processing

- **Supported FFmpeg Features**:
  - Multiple input/output streams
  - Format-specific optimization flags
  - Metadata preservation and manipulation
  - Time-based seeking and trimming
  - Custom encoding parameters

### Multi-Source Ingestion

- Local filesystem (LFS)
- HTTP/HTTPS URLs
- AWS S3 buckets
- Google Cloud Storage
- Git repositories

### Distributed Architecture

- gRPC-based service communication
- MongoDB for job persistence and state tracking
- Redis for distributed coordination and locking
- Horizontal scaling of processing workers

## Architecture

```text
┌─────────────────┐    ┌─────────────────┐
│   Client App    │───▶│   MIMO Server   │
└─────────────────┘    └─────────────────┘
                                │
                       ┌─────────────────┐
                       │   Mops Service  │
                       └─────────────────┘
                                │
                       ┌─────────────────┐
                       │    MongoDB      │
                       └─────────────────┘
                                │
        ┌───────────────────────┼───────────────────────┐
        ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│ IngestSplitter  │    │   AtomRacer     │    │   AtomPuller    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                       ┌─────────────────┐
                       │ FFmpeg Wrapper  │
                       │  - ffprobe      │
                       │  - ffmpeg       │  
                       └─────────────────┘
```

### Core Components

- **MIMO Server**: gRPC API server handling client requests
- **Mops Service**: Media operations service managing job lifecycle  
- **IngestSplitter**: Breaks ingestion requests into atomic processing tasks
- **AtomRacer**: Coordinates and monitors atom processing workers
- **AtomPuller**: Downloads media from remote sources to local storage

## Getting Started

### Prerequisites

- Go 1.19+
- FFmpeg (for media processing)
- MongoDB (for job persistence)
- Redis (for distributed coordination)

### Configuration

Create a `muxfarm.yaml` configuration file:

```yaml
Muxfarm:
  MIMO:
    HostName: "localhost"
    Port: "50050"

DocumentStore: "mongodb"
MongoDB:
  Name: "mongodb"
  URI: "mongodb://localhost:27017"
  DBName: "muxdb"
DLMRedis:
  Name: "redis" 
  URI: "localhost:6379"
  DBNumber: 11
```

### Running the System

1. Start the MIMO server:

```bash
./muxfarm mimo
```

2. Start processing workers:

```bash
./muxfarm stitch ingestSplitter
./muxfarm stitch atomRacer
```

## API Usage

MuxFarm exposes a gRPC API for submitting media processing jobs:

```go
// Submit a media processing job
mediain := &plumber.MediaIn{
    Input: []*plumber.Media{
        {
            Storagetype: plumber.StorageType_STORAGE_HTTP,
            Uri: "https://example.com/video.mp4",
        },
    },
    Operation: &plumber.Operation{
        // Define processing operations
    },
}

// Returns a job ID for tracking
jobID, err := client.Ingest(ctx, mediain)
```

## Output Formats

MuxFarm generates streaming-optimized output:

- **MPEG-DASH**: `.mpd` manifests with segmented media
- **HLS**: `.m3u8` playlists with transport segments
- **Fragmented MP4**: Initialization and media segments
- **Multi-bitrate**: Adaptive streaming support

## Development

### Protocol Buffers

The system uses Protocol Buffers for service definitions:

- `plumber/plumber.proto`: Core service API
- `fixtures/fixtures.proto`: Data persistence models
- `plumber/state.proto`: Job state management

Regenerate protobuf code after changes:

```bash
make proto  # or protoc commands
```

### Adding Storage Backends

Implement the storage interface in `plumber/plumber.go` and add the new type to `StorageType` enum.

## Monitoring

The system tracks job states through MongoDB collections:

- `ingestCollection`: Top-level ingestion jobs
- `atomCollection`: Individual processing tasks

Job states progress through: `STATE_UNSPECIFIED` → `PULL_OK` → `INGEST_OK`

## Contributing

MuxFarm is designed to be extensible. Key areas for contribution:

- Additional codec support
- New storage backend integrations  
- Enhanced monitoring and observability
- Performance optimizations

## License

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.