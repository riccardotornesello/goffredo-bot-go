<div class="modal fade" id="soundModal" tabindex="-1" aria-labelledby="soundModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <form method="POST" enctype="multipart/form-data">
          <div class="modal-header">
            <h5 class="modal-title" id="soundModalLabel">Add sound</h5>
          </div>
          <div class="modal-body">
            <div class="mb-3">
                <label for="soundName" class="form-label">Sound name</label>
                <input type="text" id="soundName" name="name" class="form-control" aria-describedby="soundHelpBlock">
                <small id="soundHelpBlock" class="form-text text-muted">Just a name to remember the sound</small>
            </div>
            <div class="mb-3">
                <label class="form-label" for="soundFile">File</label>
                <input type="file" name="sound" class="form-control" id="soundFile" accept=".mp3">
            </div>
          </div>
          <div class="modal-footer justify-content-between">
            <button type="button" class="btn bg-gradient-dark" data-bs-dismiss="modal">Close</button>
            <button type="submit" class="btn bg-gradient-primary">Add</button>
          </div>
        </form>
      </div>
    </div>
</div>