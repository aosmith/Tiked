#include <mmsystem.h>

const int NUMPTS = 44100 * 10;   // 10 seconds
int sampleRate = 44100;
short int waveIn[NUMPTS];   // 'short int' is a 16-bit type; I request 16-bit samples below
                            // for 8-bit capture, you'd use 'unsigned char' or 'BYTE' 8-bit types

HWAVEIN      hWaveIn;
WAVEHDR      WaveInHdr;
MMRESULT result;

// Specify recording parameters
WAVEFORMATEX pFormat;
pFormat.wFormatTag=WAVE_FORMAT_PCM;     // simple, uncompressed format
pFormat.nChannels=1;                    //  1=mono, 2=stereo
pFormat.nSamplesPerSec=sampleRate;      // 44100
pFormat.nAvgBytesPerSec=sampleRate*2;   // = nSamplesPerSec * n.Channels * wBitsPerSample/8
pFormat.nBlockAlign=2;                  // = n.Channels * wBitsPerSample/8
pFormat.wBitsPerSample=16;              //  16 for high quality, 8 for telephone-grade
pFormat.cbSize=0;

result = waveInOpen(&hWaveIn, WAVE_MAPPER,&pFormat,
           0L, 0L, WAVE_FORMAT_DIRECT);
if (result)
{
 char fault[256];
 waveInGetErrorText(result, fault, 256);
 Application->MessageBox(fault, "Failed to open waveform input device.",
             MB_OK | MB_ICONEXCLAMATION);
 return;
}

// Set up and prepare header for input
WaveInHdr.lpData = (LPSTR)waveIn;
WaveInHdr.dwBufferLength = NUMPTS*2;
WaveInHdr.dwBytesRecorded=0;
WaveInHdr.dwUser = 0L;
WaveInHdr.dwFlags = 0L;
WaveInHdr.dwLoops = 0L;
waveInPrepareHeader(hWaveIn, &WaveInHdr, sizeof(WAVEHDR));

// Insert a wave input buffer
result = waveInAddBuffer(hWaveIn, &WaveInHdr, sizeof(WAVEHDR));
if (result)
{
 MessageBox(Application->Handle, "Failed to read block from device",
                  NULL, MB_OK | MB_ICONEXCLAMATION);
 return;
}


// Commence sampling input
result = waveInStart(hWaveIn);
if (result)
{
 MessageBox(Application->Handle, "Failed to start recording",
                  NULL, MB_OK | MB_ICONEXCLAMATION);
 return;
}


// Wait until finished recording
do {} while (waveInUnprepareHeader(hWaveIn, &WaveInHdr, sizeof(WAVEHDR))==WAVERR_STILLPLAYING);

waveInClose(hWaveIn);