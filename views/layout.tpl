<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

  <!-- Favicon -->
  <link rel="apple-touch-icon" sizes="180x180" href="/static/ico/apple-touch-icon.png">
  <link rel="icon" type="image/png" sizes="32x32" href="/static/ico/favicon-32x32.png">
  <link rel="icon" type="image/png" sizes="16x16" href="/static/ico/favicon-16x16.png">
  <link rel="manifest" href="/static/ico/site.webmanifest">
  <link rel="mask-icon" href="/static/ico/safari-pinned-tab.svg" color="#5bbad5">
  <link rel="shortcut icon" href="/static/ico/favicon.ico">
  <meta name="msapplication-TileColor" content="#da532c">
  <meta name="msapplication-config" content="/static/ico/browserconfig.xml">
  <meta name="theme-color" content="#ffffff">

  <title>
    Goffredo Bot
  </title>
  
  <!--     Fonts and icons     -->
  <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,600,700" rel="stylesheet" />
  <!-- Nucleo Icons -->
  <link href="/static/css/nucleo-icons.css" rel="stylesheet" />
  <link href="/static/css/nucleo-svg.css" rel="stylesheet" />
  <!-- Font Awesome Icons -->
  <script src="https://kit.fontawesome.com/42d5adcbca.js" crossorigin="anonymous"></script>
  <!-- CSS Files -->
  <link id="pagestyle" href="/static/css/soft-ui-dashboard.css?v=1.0.3" rel="stylesheet" />

  <!-- Global site tag (gtag.js) - Google Analytics -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=G-H3V4X0DWM9"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());

    gtag('config', 'G-H3V4X0DWM9');
  </script>
</head>

<body>
  <div class="container-fluid wrapper">
    <main class="main-content position-relative max-height-vh-100 h-100 mt-1 border-radius-lg ">

      {{.LayoutContent}}

    </main>

    {{template "partials/footer.tpl" .}}
  </div>

  <!--   Core JS Files   -->
  <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js" integrity="sha384-IQsoLXl5PILFhosVNubq5LC7Qb9DXgDA9i+tQ8Zj3iwWAwPtgFTxbJ8NT4GN1R8p" crossorigin="anonymous"></script>
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.min.js" integrity="sha384-cVKIPhGWiC2Al4u+LWgxfKTRIcfu0JTxR+EQDz/bgldoEyl4H0zUF0QKbrJ0EcQF" crossorigin="anonymous"></script>
  <script src="/static/js/plugins/perfect-scrollbar.min.js"></script>
  <script src="/static/js/plugins/smooth-scrollbar.min.js"></script>

  <script>
    var win = navigator.platform.indexOf('Win') > -1;
    if (win && document.querySelector('#sidenav-scrollbar')) {
      var options = {
        damping: '0.5'
      }
      Scrollbar.init(document.querySelector('#sidenav-scrollbar'), options);
    }
  </script>

  <!-- Github buttons -->
  <script async defer src="https://buttons.github.io/buttons.js"></script>

  <!-- Control Center for Soft Dashboard: parallax effects, scripts for the example pages etc -->
  <script src="/static/js/soft-ui-dashboard.min.js?v=1.0.3"></script>

  <script src="/static/js/goffredo.js"></script>
</body>
</html>
