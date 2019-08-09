//
//  ViewController.swift
//  ratara
//
//  Created by 澤田　雄一 on 2019/08/06.
//  Copyright © 2019 澤田　雄一. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    @IBOutlet var food_place: UILabel!
    @IBOutlet var food_country: UILabel!
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
        food_place.text = "ばしょ！"
        food_country.text = "国"
    }


}

