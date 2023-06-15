let systemPkgs = import <nixpkgs> {};

in {
	# gotk4-nix = ../../gotk4-nix;
	gotk4-nix = systemPkgs.fetchFromGitHub {
		owner = "brotholo";
		repo = "gotk4-nix";
		# rev = "b186ac5419c22f4b75b1bdc35ef4fd9349c6b65e";
        # rev = "a0d85f906d67354dee3d9c8456ef41e67c044f49";
        rev = "8f007e76a0d87a57a79a3b4403ce8a8e15d4a424";
        sha256 = "sha256-B0tvAIz88nT5ZZMaln8GIHhnSe5EssR5HPgbwLpZPRM=";
        # a0d85f906d67354dee3d9c8456ef41e67c044f49
        # sha256 = "sha256-Q6kq+vJB9Om7DdUUB8plFOdmIulLM9bvUqF7Krm1nvI=";
		# sha256 = "1pfx0p4w56d7pa1l9ipqfq52znfl7slc2wbjfqddq1jv1fp9z43q";
	};
}
