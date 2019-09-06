package com.example.ratara

import android.os.Bundle
import android.support.design.widget.Snackbar
import android.support.v7.app.AppCompatActivity;
import android.view.Menu
import android.view.MenuItem
import android.view.View
import android.widget.Button
import android.widget.TextView

import kotlinx.android.synthetic.main.activity_main.*
import ratara.Ratara
import android.widget.Toast



class MainActivity : AppCompatActivity() {
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_main)
            setSupportActionBar(toolbar)
            val placeView = findViewById<TextView>(R.id.food_place) as TextView
            val countryView = findViewById<TextView>(R.id.food_country) as TextView
            val randButton = findViewById<Button>(R.id.rand_button) as Button

            placeView.text = ""
            countryView.text = ""

            randButton.setOnClickListener { _ ->
                val place = Ratara.place()
                val country = Ratara.foodCountry()
                placeView.text = place
                countryView.text = country
            }

            fab.setOnClickListener { view ->
                Snackbar.make(view, "Replace with your own action", Snackbar.LENGTH_LONG)
                    .setAction("Action", null).show()
            }
        }

    override fun onCreateOptionsMenu(menu: Menu): Boolean {
        // Inflate the menu; this adds items to the action bar if it is present.
        menuInflater.inflate(R.menu.menu_main, menu)
        return true
    }

    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        return when (item.itemId) {
            R.id.action_settings -> true
            else -> super.onOptionsItemSelected(item)
        }
    }
}
