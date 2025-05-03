document.getElementById('startTestButton').addEventListener('click', async () => {
    const startTime = performance.now();

    try {
        const response = await fetch('/download');
        if (!response.ok) {
            throw new Error('Failed to download file');
        }

        const contentLength = response.headers.get('Content-Length');
        const totalBytes = contentLength ? Number.parseInt(contentLength, 10) : 0;
        const fileSizeMB = totalBytes / (1024 * 1024);

        let receivedBytes = 0;
        const reader = response.body.getReader();
        const chunks = [];
        const decoder = new TextDecoder();

        while (true) {
            const { done, value } = await reader.read();
            if (done) break;

            chunks.push(value);
            receivedBytes += value.length;

            const elapsedSeconds = (performance.now() - startTime) / 1000;
            const speedMbps = (receivedBytes / (1024 * 1024) / elapsedSeconds) * 8;

            document.getElementById('timeTaken').innerText = `Time Taken: ${elapsedSeconds.toFixed(2)} seconds`;
            document.getElementById('speedMbps').innerText = `Speed: ${speedMbps.toFixed(2)} Mbps`;
            document.getElementById('filesize').innerText = `Downloaded: ${(receivedBytes / (1024 * 1024)).toFixed(2)} MB of ${fileSizeMB.toFixed(2)} MB`;
        }

        const blob = new Blob(chunks);
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'downloadFile.bin';
        document.body.appendChild(a);
        a.click();
        a.remove();
        window.URL.revokeObjectURL(url);

        const endTime = performance.now();
        const durationSeconds = (endTime - startTime) / 1000;
        const finalSpeedMbps = (fileSizeMB / durationSeconds) * 8;

        document.getElementById('timeTaken').innerText = `Time Taken: ${durationSeconds.toFixed(2)} seconds`;
        document.getElementById('speedMbps').innerText = `Speed: ${finalSpeedMbps.toFixed(2)} Mbps`;
        document.getElementById('filesize').innerText = `File Size: ${fileSizeMB.toFixed(2)} MB`;
    } catch (error) {
        console.error('Error during download:', error);
        alert('Failed to download the file.');
    }
});