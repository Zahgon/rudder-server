package batch

// This is going to call appropriate method of Filemanager & DeleteManager
// to get deletion done.
// called by delete/deleteSvc with (model.Job, model.Destination).
// returns final status,error ({successful, failure}, err)
import (
	"context"
	"sync"

	_ "go.uber.org/automaxprocs"

	"github.com/rudderlabs/rudder-go-kit/filemanager"
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/delete/batch/filehandler"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
)

var (
	pkgLogger             = logger.NewLogger().Child("batch")
	StatusTrackerFileName = "rudderDeleteTracker.txt"
	supportedDestinations = []string{"S3", "S3_DATALAKE"}
)

type Batch struct {
	mu         sync.Mutex
	FM         filemanager.FileManager
	session    filemanager.ListSession
	TmpDirPath string
}

// listFiles fetches the files from filemanager under prefix mentioned and for a
// specified limit.
func (b *Batch) listFiles(ctx context.Context, prefix string, limit int) (fileObjects []*filemanager.FileInfo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// two pointer algorithm implementation to remove all the files from which users are already deleted.
func removeCleanedFiles(files []*filemanager.FileInfo, cleanedFiles []string) []*filemanager.FileInfo {
	_ = "STUB: not implemented"
	return nil
}

// append <fileName> to <statusTrackerFile> locally for which deletion has completed.
// updateStatusTrackerFile updates the tracker file with the fileName information
func (*Batch) updateStatusTrackerFile(absStatusTrackerFileName, fileName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Batch) cleanedFiles(_ context.Context, path string, job *model.Job) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if statusTracker.txt exists then read it & remove all those files name from above gzFilesObjects,
// since those files are already cleaned.

// insert <jobID> in 1st line

// check if our <jobID> matches with the one in file.
// if not, then truncate the file & write new current jobID.

// This might happen when we have a job partially working on the
// suppress with delete and then it fails and second job starts in the meantime.
// So we keep the latest state in here.

// truncate the contents of the file, to start writing for another
// <jobID> information.

// if we have entries then read it.

// downloads `fileName` locally. And returns empty file, if file not found.
// Note: download happens concurrently in 5 go routine by default.
func (b *Batch) download(ctx context.Context, completeFileName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func downloadWithExpBackoff(ctx context.Context, fu func(context.Context, string) (string, error), fileName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func uploadWithExpBackoff(ctx context.Context, fu func(ctx context.Context, uploadFileAbsPath, actualFileName, absStatusTrackerFileName string) error, uploadFileAbsPath, actualFileName, absStatusTrackerFileName string) error {
	_ = "STUB: not implemented"
	return nil
}

// replace old json.gz & statusTrackerFile with the new during upload.
// Note: upload happens concurrently in 5 go routine by default
func (b *Batch) upload(_ context.Context, uploadFileAbsPath, actualFileName, absStatusTrackerFileName string) error {
	_ = "STUB: not implemented"
	return nil
}

type BatchManager struct {
	FilesLimit int
	FMFactory  filemanager.Factory
}

func (*BatchManager) GetSupportedDestinations() []string { _ = "STUB: not implemented"; return nil }

// Delete users corresponding to input userAttributes from a given batch destination
func (bm *BatchManager) Delete(
	ctx context.Context,
	job model.Job,
	destination *backendconfig.DestinationT,
) model.JobStatus {
	_ = "STUB: not implemented"
	return *new(model.JobStatus)
}

// parent directory of all the temporary files created/downloaded in the process of deletion.

// Get the prefix which should be the base of the
// of the cleanup operations.

// Get filehandler from a factory on every iteration, to not share the data.

// TODO: Why not have a common function to upload the cleaned files to tracker in one shot ?
// Why do it one entry at a time ?

func LocalFileHandlerFactory(dest, upstreamFilePath string) filehandler.LocalFileHandler {
	_ = "STUB: not implemented"
	return *new(filehandler.LocalFileHandler)
}

// handleIdentityRemoval is a convenience wrapper over the filehandler
// performing the operations over the file to remove the user identity.
func handleIdentityRemoval(
	ctx context.Context,
	handler filehandler.LocalFileHandler,
	attributes []model.User,
	sourceFile, targetFile string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func maxRoutines() int { _ = "STUB: not implemented"; return 0 }

func getFileSize(fileAbsPath string) int { _ = "STUB: not implemented"; return 0 }

func (b *Batch) cleanup(ctx context.Context, prefix string) { _ = "STUB: not implemented"; return }
