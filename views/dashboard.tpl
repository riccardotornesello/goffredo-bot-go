{{template "modals/sound.tpl"}}

  <nav class="navbar navbar-main navbar-expand-lg px-0 mx-4 shadow-none border-radius-xl" id="navbarBlur" navbar-scroll="true">
    <div class="container-fluid py-1 px-3">
      <div class="collapse navbar-collapse mt-sm-0 mt-2 me-md-0 me-sm-4" id="navbar">
        <ul class="ms-md-auto navbar-nav justify-content-end">
          <li class="nav-item d-flex align-items-center">
            <a href="/logout" class="nav-link text-body font-weight-bold px-0">
              <i class="fa fa-user me-sm-1"></i>
              <span class="d-sm-inline d-none">LOGOUT</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </nav>

  <div class="container-fluid py-4">
    <a href="{{ .BotAddUrl }}" type="button" class="btn bg-gradient-primary btn-lg">Add Goffredo in your server</a>
  </div>

  <div class="container-fluid py-4">
    <div class="row">
      <div class="col-12">
        <div class="card mb-4">
          <div class="card-header pb-0 p-3">
            <div class="row">
              <div class="col-6">
                <h6 class="mb-0">Sounds</h6>
                <p class="text-sm mb-0">
                  <i class="fa fa-tasks text-info" aria-hidden="true"></i>
                  {{ .SoundsLimit }} maximum sounds,<span class="font-weight-bold ms-1">{{ .SoundsLeft }} left</span>
                </p>
              </div>
              <div class="col-6 text-end">
                <button {{ if lt .SoundsLeft 1 }}disabled{{ end }} type="button" class="btn bg-gradient-primary mb-0" data-bs-toggle="modal" data-bs-target="#soundModal"><i class="fas fa-plus"></i>&nbsp;&nbsp;Add new sound</button>
              </div>
            </div>
          </div>
          <div class="card-body px-0 pt-0 pb-2">
            <div class="table-responsive p-0">
              <table class="table align-items-center mb-0">
                <thead>
                  <tr>
                    <th class="text-uppercase text-secondary text-xxs font-weight-bolder opacity-7">Sound</th>
                    <th class="text-center text-uppercase text-secondary text-xxs font-weight-bolder opacity-7">Status</th>
                    <th class="text-secondary opacity-7"></th>
                  </tr>
                </thead>
                <tbody>
                  {{ range .Sounds }}
                  <tr>
                    <td>
                      <div class="d-flex px-2 py-1">
                        <div>
                          <img src="/static/img/sound.png" class="avatar avatar-sm bg-gradient-primary me-3 p-2" alt="sound">
                        </div>
                        <div class="d-flex flex-column justify-content-center">
                          <h6 class="mb-0 text-sm">{{ .Name }}</h6>
                        </div>
                      </div>
                    </td>
                    <td class="align-middle text-center text-sm">
                      <span class="badge badge-sm bg-gradient-success">Enabled</span>
                    </td>
                    <td class="align-middle">
                      <a onclick="deleteSound({{ .Id }})" class="text-secondary font-weight-bold text-xs" data-toggle="tooltip" data-original-title="Delete sound">
                        Delete
                      </a>
                    </td>
                  </tr>
                  {{ end }}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>