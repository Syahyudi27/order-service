{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # Ini baris yang perlu Anda pastikan ada dan benar:
    go_1_22

    # Anda bisa menambahkan alat lain yang Anda butuhkan di sini, contoh:
    # git
    # delve # untuk debugging Go
    # nodejs # jika proyek Anda punya frontend
  ];

  # Jika Anda memiliki bagian shellHook, biarkan saja atau tambahkan jika diperlukan
  # shellHook = ''
  #   echo "Selamat datang di lingkungan Go!"
  # '';
}